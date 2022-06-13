import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Category } from '../interfaces/category';
import { CredentialControlService } from './credential-control.service';
import { Config } from '../modules/config-module'

@Injectable({
  providedIn: 'root'
})
export class CategoryCrudServiceService {
  url = Config.address + ':' + Config.port;

  urlGetCategoriesAdmin = this.url + '/api/categorias';
  urlGetCategoriesUser = this.url + '/api/categorias/padre';
  urlGetCategories = this.urlGetCategoriesAdmin;

  urlCreateCategory = this.url + '/api/categorias';
  urlDeleteCategory = this.url + '/api/categorias/';
  urlEditCategory = this.url + '/api/categorias/';
  total = 0;

  constructor(private http: HttpClient) {
    if (localStorage.getItem('Role') != undefined) {
      let rol = localStorage.getItem('Role');
      if (rol?.includes('ADMIN_ROLE')) {
        this.urlGetCategories = this.urlGetCategoriesAdmin;
      } 
      else 
      {
        this.loadTotalCategoriesBasicUser();
      }
      
    }
  }

  loadTotalCategoriesBasicUser(){
    const httpOptions = {
      headers: new HttpHeaders({
        'Authorization': 'Bearer ' + CredentialControlService.getToken()
      })
    };
    (this.http.get(this.urlGetCategoriesUser + '?limite=1&desde=0', httpOptions) as Observable<Category>).subscribe(
      res => {
        this.total = res['total'];
        this.urlGetCategories = this.urlGetCategoriesUser + '?limite=' + this.total + '&desde=0';
        console.log(res);
      },
      err => {
        console.log(err);
      }
    );
    
  }

  getCategories() {

    const fullyQualifiedUrl = this.urlGetCategories;
    console.log(fullyQualifiedUrl);


    const httpOptions = {
      headers: new HttpHeaders({
        'Authorization': 'Bearer ' + CredentialControlService.getToken()
      })
    };

    return this.http.get(this.urlGetCategories, httpOptions) as Observable<Category>;
  }

  createCategory(nombre: string) {
    const httpOptions = {
      headers: new HttpHeaders({
        'Authorization': 'Bearer ' + CredentialControlService.getToken(),
        "Content-Type": "application/json"
      })
    };

    const body = {
      'nombre': nombre
    }

    return this.http.post(this.urlCreateCategory, body, httpOptions);
  }

  deleteCategory(id: string) {

    const httpOptions = {
      headers: new HttpHeaders({
        'Authorization': 'Bearer ' + CredentialControlService.getToken(),
      })
    };

    return this.http.delete(this.urlDeleteCategory + id, httpOptions);
  }

  editCategory(id: string, nombre: string) {
    const httpOptions = {
      headers: new HttpHeaders({
        'Authorization': 'Bearer ' + CredentialControlService.getToken(),
        "Content-Type": "application/json"
      })
    };

    const body = {
      'nombre': nombre
    }

    return this.http.put(this.urlEditCategory + id, body, httpOptions);
  }
}
