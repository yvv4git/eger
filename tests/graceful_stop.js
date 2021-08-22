import http from 'k6/http';

export let options = {
    discardResponseBodies: true,
    scenarios: {
        contacts: {
            executor: 'constant-vus',
            vus: 100,
            duration: '10s',
            gracefulStop: '3s',
        },
    },
};

export default function () {
    let delay = Math.floor(Math.random() * 5) + 1;
    http.get(`https://httpbin.test.k6.io/delay/${delay}`);
}
