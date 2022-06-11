import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { LoginInterface } from '../interfaces/login-interface';
import { GetVideoInterface } from '../interfaces/get-video-interface';
import { Config } from '../modules/config-module'

@Injectable({
  providedIn: 'root'
})

export class CredentialControlService {

  url = Config.address + ':' + Config.port;
  urlLogin = this.url + '/api/auth/login';
  urlGetVideosAdmin = this.url + '/api/videos?limite=1&desde=0';
  urlGetVideosUser = this.url + '/api/videos/padre?limite=1&desde=0';
  urlFull = this.urlGetVideosAdmin;

  constructor(private http: HttpClient) { 
    if(localStorage.getItem('Role') != undefined){
      let rol = localStorage.getItem('Role');
      if(rol?.includes('ADMIN_ROLE'))
        this.urlFull = this.urlGetVideosAdmin;
      else
        this.urlFull = this.urlGetVideosUser;
    }
  }

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

    return this.http.get(this.urlFull, httpOptions) as Observable<GetVideoInterface>;
  }

  logOut(){
    localStorage.clear();
    location.assign("/Login");
  }
}
