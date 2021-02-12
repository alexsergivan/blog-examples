import { sleep } from 'k6';
import http from 'k6/http';
import { check } from 'k6';
import { Rate } from 'k6/metrics';

export let errorRate = new Rate('errors');

export let options = {
    scenarios: {
        sign_in_page_test: {
            executor: 'constant-vus',
            duration: '1m',
            vus: 100,
            tags: { test_type: 'signInPage' },
            exec: 'signInPage',
        },
    },

    thresholds: {
        http_req_duration: ['avg<500'], // avg response times must be below 0.5s
        errors: ['rate<0.1'], // <10% errors
    },
};

export function signInPage() {
    const res = http.get(getDomain() + '/user/signin');
    const result = check(res, {
        'status is 200': (r) => r.status == 200,
    });
    errorRate.add(!result);
    sleep(3)
}

function getDomain() {
    return __ENV.DOMAIN;
}
