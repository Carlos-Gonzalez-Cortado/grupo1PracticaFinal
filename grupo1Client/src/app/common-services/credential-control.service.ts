import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { LoginInterface } from '../interfaces/login-interface';
import { GetVideoInterface } from '../interfaces/get-video-interface';

@Injectable({
  providedIn: 'root'
})

export class CredentialControlService {

  url = 'https://labinfsoft.herokuapp.com';
  urlLogin = this.url + '/api/auth/login';
  urlGetVideos = this.url + '/api/videos?limite=1&desde=0';

  constructor(private http: HttpClient) { }

  sendCredential(email: string, password: string) {

    const body_content = {
      "correo": email,
      "password": password
    }

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      })
    };

    return this.http.post(this.urlLogin, body_content, httpOptions) as Observable<LoginInterface>;
  }

  static getToken() {
    return localStorage.getItem('Token');
  }

  checkCredential() {

    const body_content = {
      "token": localStorage.getItem("Token")
    }

    const httpOptions = {
      headers: new HttpHeaders({
        'Authorization': 'Bearer ' + CredentialControlService.getToken()
      })
    };

    console.log(httpOptions);


    return this.http.get(this.urlGetVideos, httpOptions) as Observable<GetVideoInterface>;
  }
}
