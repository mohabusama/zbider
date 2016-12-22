import { Component } from '@angular/core';
import { OnInit } from '@angular/core';

import { Result } from './result.model'

import { SearchService } from './search.service'


export class User {
    id: number;
    name: string;
}

@Component({
    selector: 'my-app',
    providers: [SearchService],
    template: `
        <md-toolbar [color]="myColor">
            <span></span>

            <md-input [(ngModel)]="query"></md-input>
            <button md-button (click)="Search()"><md-icon>search</md-icon></button>

        </md-toolbar>
        <md-card *ngIf="result.results.length === 0">
            <md-card-content><h3>No search results</h3></md-card-content>
        </md-card>
        <md-card *ngFor="let res of result.results">
            <md-card-title>
                <a href="{{res.data.zbider_fields.link}}" target="_blank">
                {{res.data.zbider_fields.title}}
                </a>
            </md-card-title>
            <md-card-content>
                <p>
                    {{res.data.zbider_fields.text}}
                </p>
            </md-card-content>
        </md-card>
   `,
})
export class AppComponent  {
    query = '';
    user: User = {
        id: 1,
        name: "ZBIDER USER"
    };

    result: Result = {
        total_hits: 0,
        took: 0.0,
        results: []
    };

    constructor(private searchService: SearchService) {}

    ngOnInit(): void {}

    Search(): void {
        this
            .searchService
            .search(this.query)
            .then((res) => {
                console.log(res);
                this.result = res;
            });
    }
}
