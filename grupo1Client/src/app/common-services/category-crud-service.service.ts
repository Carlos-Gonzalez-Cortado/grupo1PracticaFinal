import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Category } from '../interfaces/category';
import { CredentialControlService } from './credential-control.service';

@Injectable({
  providedIn: 'root'
})
export class CategoryCrudServiceService {
  url = 'https://labinfsoft.herokuapp.com';
  urlGetCategories = this.url + '/api/categorias';

  constructor(private http: HttpClient) { }

  getCategories(){

    const fullyQualifiedUrl = this.urlGetCategories;

    console.log(fullyQualifiedUrl);
    

    const httpOptions = {
      headers: new HttpHeaders({
        'Authorization': 'Bearer ' + CredentialControlService.getToken()
      })
    };

    return this.http.get(this.urlGetCategories, httpOptions) as Observable<Category>;
  }
}
