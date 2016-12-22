import { Injectable } from '@angular/core';
import { Http, RequestOptionsArgs, URLSearchParams } from '@angular/http';

import 'rxjs/add/operator/toPromise';

import { Result } from './result.model';

@Injectable()
export class SearchService {
    private searchUrl = 'search';

    constructor(private http: Http) {}

    search(query: string): Promise<Result> {
        let q = new URLSearchParams();
        q.set('q', query)

        return this
                .http
                .get(this.searchUrl, {search: q})
                // .get(this.searchUrl)
                .toPromise()
                .then(response => response.json() as Result)
                .catch((err: any): Promise<any> => {
                    console.log(err);
                    return Promise.reject(err.message || err);
                });
    }
}
