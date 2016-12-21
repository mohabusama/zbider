import { Component } from '@angular/core';

export class User {
    id: number;
    name: string;
}

export class Result {
    score: number;
    zbider_type: string;
    data: Object
}

const RESULTS: Result[] = [
    {score: 0.1, zbider_type: 'zbider-user', data: {link: 'https://www.google.de', title: 'T 0.1', description: 'lorem'}},
    {score: 0.12, zbider_type: 'zbider-user', data: {link: 'https://www.google.de', title: 'T 0.12', description: 'lorem'}}
]

@Component({
    selector: 'my-app',
    template: `
        <md-toolbar [color]="myColor">
            <span>ZBIDER</span>

            <md-input [(ngModel)]="query"></md-input>
            <button md-button (click)="Search()"><md-icon>search</md-icon></button>

        </md-toolbar>
        <md-card *ngIf="results.length === 0">
            <md-card-content>No search results</md-card-content>
        </md-card>
        <md-card *ngFor="let res of results">
            <md-card-title>{{res.data.title}}</md-card-title>
            <md-card-content>
                <p>
                    {{res.data.description}}
                </p>
            </md-card-content>
        </md-card>
        <md-input [(ngModel)]="user.name"></md-input>
   `,
})
export class AppComponent  {
    query = '';
    user: User = {
        id: 1,
        name: "ZBIDER USER"
    };

    results = RESULTS;
    // results = [];

    Search(): void {
        console.log(this.query);
    }
}
