import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {RouterModule, Routes} from "@angular/router";
import {LoginPageComponent} from "../components/login-page/login-page.component";
import {CataloguePageComponent} from "../components/catalogue-page/catalogue-page.component";
import {RecommendationsPageComponent} from "../components/recommendations-page/recommendations-page.component";

const routes: Routes = [
  { path: 'home', component: CataloguePageComponent},
  { path: 'recommendations', component: RecommendationsPageComponent},
  { path: 'login', component: LoginPageComponent},
  { path: '**', redirectTo: 'home'},
];

@NgModule({
  declarations: [],
  // imports: [
  //   CommonModule
  // ]
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
