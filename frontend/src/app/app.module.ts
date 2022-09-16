import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { TopbarComponent } from './components/topbar/topbar.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {RouterLink} from "@angular/router";
import {MatIconModule} from "@angular/material/icon";
import { LoginPageComponent } from './components/login-page/login-page.component';
import {AppRoutingModule} from "./app-routing/app-routing.module";
import { CataloguePageComponent } from './components/catalogue-page/catalogue-page.component';
import { RecommendationsPageComponent } from './components/recommendations-page/recommendations-page.component';
import {FormsModule, ReactiveFormsModule} from "@angular/forms";

@NgModule({
  declarations: [
    AppComponent,
    TopbarComponent,
    LoginPageComponent,
    CataloguePageComponent,
    RecommendationsPageComponent
  ],
  imports: [
    AppRoutingModule,
    BrowserModule,
    BrowserAnimationsModule,
    RouterLink,
    MatIconModule,
    FormsModule,
    ReactiveFormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
