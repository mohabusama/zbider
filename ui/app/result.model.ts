// export const RESULTS: Result[] = [
//     {score: 0.1, zbider_type: 'zbider-user', data: {link: 'https://www.google.de', title: 'T 0.1', description: 'lorem'}},
//     {score: 0.12, zbider_type: 'zbider-user', data: {link: 'https://www.google.de', title: 'T 0.12', description: 'lorem'}}
// ]

export class Result {
    total_hits: number;
    took: number;
    results: Object[]
}
