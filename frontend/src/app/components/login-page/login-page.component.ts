import {Component, OnInit} from '@angular/core';
import {FormControl, FormGroup, Validators} from "@angular/forms";
import {AuthService} from "../../services/auth.service";
import {Router} from "@angular/router";

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent implements OnInit {
  form: FormGroup;

  constructor(private authService: AuthService, private router: Router) {
  }

  public error: string;

  ngOnInit(): void {
    this.form = new FormGroup({
      login: new FormControl(null, [Validators.required]),
      password: new FormControl(null, [Validators.required])
    })
  }

  onLogin() {
    this.authService.login(this.form.value.username, this.form.value.password).subscribe(
      session => {
        localStorage.setItem('user-token', session.token);
        localStorage.setItem('username', session.username);
        localStorage.setItem('id', session.id.toString());
        this.router.navigate(['/home']);
      },
      error => {
        if (error.status === 401) {
          this.error = 'Неправильные данные.';
        } else if (error.status === 400) {
          this.error = 'Пожалуйста, заполните все поля';
        } else {
          this.error = 'Ошибка на cервере';
        }
      }
    );
  }

  onRegister() {
    this.authService.register(this.form.value.username, this.form.value.password).subscribe(
      newuser => {
        this.authService.login(this.form.value.username, this.form.value.password).subscribe(
          session => {
            localStorage.setItem('user-token', session.token);
            localStorage.setItem('username', session.username);
            localStorage.setItem('id', session.id.toString());
            this.router.navigate(['/home']);
          },
          error => console.log(error)
        );
      },
      error => {
        if (error.status === 401) {
          this.error = 'Неправильные данные.';
        } else if (error.status === 400) {
          if (error.error.detail === 'invalid_email') {
            this.error = 'Пользователь с таким email уже существует.';
          } else if (error.error.detail === 'invalid_username') {
            this.error = 'Пользователь с таким именем уже существует.';
          } else {
            this.error = 'Произошла ошибка, проверьте, что все поля заполнены верно.';
          }
        } else {
          this.error = 'Ошибка на cервере';
        }
      }
    );
  }
}
