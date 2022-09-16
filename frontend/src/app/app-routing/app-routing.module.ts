import {NgModule} from '@angular/core';
import {RouterModule, Routes} from "@angular/router";
import {LoginPageComponent} from "../components/login-page/login-page.component";
import {CataloguePageComponent} from "../components/catalogue-page/catalogue-page.component";
import {RecommendationsPageComponent} from "../components/recommendations-page/recommendations-page.component";

const routes: Routes = [
  {path: 'catalogue', component: CataloguePageComponent},
  {path: 'recommendations', component: RecommendationsPageComponent},
  {path: 'login', component: LoginPageComponent},
  {path: 'home', redirectTo: 'catalogue'},
  {path: '**', redirectTo: 'catalogue'},
];

@NgModule({
  declarations: [],
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
