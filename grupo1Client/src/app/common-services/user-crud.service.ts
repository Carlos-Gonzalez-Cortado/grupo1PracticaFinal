import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Users } from '../interfaces/users';
import { CredentialControlService } from './credential-control.service';
import { Config } from '../modules/config-module'

@Injectable({
  providedIn: 'root'
})
export class UserCrudService {
  
  url = Config.address + ':' + Config.port;
  urlGetUsers = this.url + "/api/usuarios";
  urlCreateUser = this.url + "/api/usuarios";
  urlEditUser = this.url + "/api/usuarios/";
  urlDeleteUser = this.url + "/api/usuarios/";

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

  createUser(nombre: string, correo: string, password: string) {
    const httpOptions = {
      headers: new HttpHeaders({
        'Authorization': 'Bearer ' + CredentialControlService.getToken(),
        "Content-Type": "application/json"
      })
    };

    const body = {
      'nombre': nombre,
      'correo': correo,
      'password': password
    }

    return this.http.post(this.urlCreateUser, body, httpOptions);
  }

  deleteUser(id: string) {

    const httpOptions = {
      headers: new HttpHeaders({
        'Authorization': 'Bearer ' + CredentialControlService.getToken(),
      })
    };

    return this.http.delete(this.urlDeleteUser + id, httpOptions);
  }

  editUser(id: string, nombre: string, correo: string, password: string) {
    const httpOptions = {
      headers: new HttpHeaders({
        'Authorization': 'Bearer ' + CredentialControlService.getToken(),
        "Content-Type": "application/json"
      })
    };

    const body = {
      'nombre': nombre,
      'correo': correo,
      'password': password
    }    

    return this.http.put(this.urlEditUser + id, body, httpOptions);
  }

}
