import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Users } from '../interfaces/users';
import { CredentialControlService } from './credential-control.service';

@Injectable({
  providedIn: 'root'
})
export class UserCrudService {
  
  url = 'https://labinfsoft.herokuapp.com';
  urlGetUsers = this.url + "/api/usuarios";

  constructor(private http: HttpClient) { }

  getUsers(){
    const fullyQualifiedUrl = this.urlGetUsers;

    console.log(fullyQualifiedUrl);
    
    const httpOptions = {
      headers: new HttpHeaders({
        'Authorization': 'Bearer ' + CredentialControlService.getToken()
      })
    };

    return this.http.get(this.urlGetUsers, httpOptions) as Observable<Users>;
  }

  

}
