import { NgModule }      from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { MaterialModule } from '@angular/material';
import { FormsModule }   from '@angular/forms';
import { HttpModule }    from '@angular/http';

// import { InMemoryWebApiModule } from 'angular-in-memory-web-api';
// import { InMemoryDataService }  from './data.service';

import { AppComponent }  from './app.component';

@NgModule({
    imports:[
        BrowserModule,
        MaterialModule.forRoot(),
        // InMemoryWebApiModule.forRoot(InMemoryDataService),
        HttpModule,
        FormsModule,
    ],
    declarations: [ AppComponent ],
    bootstrap:    [ AppComponent ]
})
export class AppModule { }
