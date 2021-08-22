import http from 'k6/http';
import { sleep, check } from 'k6';
import { Rate } from 'k6/metrics';

export let options = {
    stages: [
        { duration: "5s", target: 5 }, // поднимаем количество пользователей с 1 до 5 в течении 5 секунд
        { duration: "10s", target: 5 }, // количество одновременных пользователей остается 5 в течении следующих 10 секунд
        { duration: "5s", target: 0 }   // понижаем количество одновременных пользователей с 5 до 0 в течении оставшихся 5 секунд
    ],
    thresholds: {
        errors: ['rate<0.1'], // <10% errors
    },
};

export let errorRate = new Rate('errors');


export default function () {
    let res = http.get('http://10.0.2.15:8080/api/v1/jaeger');
    check(res, {
        'is status 200': (r) => r.status === 200,
    });

    //sleep(1);
    errorRate.add(!res);
}