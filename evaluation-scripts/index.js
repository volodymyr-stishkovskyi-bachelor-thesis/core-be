import http from 'k6/http';
import { sleep } from 'k6';

export let options = { vus: 5, duration: '120s' };

export default function () {
    const payload = JSON.stringify({
        query: 'List my cloud-related certifications',
        chat: `${__VU}-${__ITER}`
    });
    http.post('http://localhost:8080/queries', payload, {
        headers: { 'Content-Type': 'application/json' }
    });
    sleep(1);
}