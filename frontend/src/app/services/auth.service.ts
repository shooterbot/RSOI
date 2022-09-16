import {Injectable} from "@angular/core";
import {Observable} from "rxjs";
import {HttpClient} from "@angular/common/http";
import {environment} from "../../environments/environment";
import {Session} from "../models/session";
import {User} from "../models/user";

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private url = `${environment.baseUrl}:${environment.backendPort}/api/v1`;

  constructor(private http: HttpClient) {
  }

  register(username: string, password: string): Observable<User> {
    const source = {
      username,
      password,
    };
    return this.http.post<User>(`${this.url}/users/`, source);
  }

  public login(username: string, password: string): Observable<Session> {
    const source = {
      username,
      password,
    };
    return this.http.post<Session>(`${this.url}/sessions/`, source);
  }
}
