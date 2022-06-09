import { HttpClient, HttpHeaders, HttpRequest, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { GetVideoInterface } from '../interfaces/get-video-interface';
import { CredentialControlService } from './credential-control.service';

@Injectable({
  providedIn: 'root'
})
export class VideoCrudServiceService {

  url = 'https://labinfsoft.herokuapp.com';

  urlGetVideosAdmin = this.url + '/api/videos';
  urlGetVideosUser = this.url + '/api/videos/padre';
  urlGetVideos = this.urlGetVideosAdmin;

  urlCreateVideo = this.url + '/api/videos';
  urlDeleteVideo = this.url + '/api/videos/';
  urlEditVideo = this.url + '/api/videos/';


  constructor(private http: HttpClient) { 
    if(localStorage.getItem('Role') != undefined){
      let rol = localStorage.getItem('Role');
      if(rol?.includes('ADMIN_ROLE'))
        this.urlGetVideos = this.urlGetVideosAdmin;
      else
        this.urlGetVideos = this.urlGetVideosUser;
    }
  }

  getVideos(limite:number, desde:number){

    const fullyQualifiedUrl = this.urlGetVideos + '?limite=' + limite + '&desde=' + desde;

    console.log(fullyQualifiedUrl);
    

    const httpOptions = {
      headers: new HttpHeaders({
        'Authorization': 'Bearer ' + CredentialControlService.getToken()
      })
    };

    return this.http.get(fullyQualifiedUrl, httpOptions) as Observable<GetVideoInterface>;
  }

  createVideo(nombre:string, url:string, categoria:string){
    const httpOptions = {
      headers: new HttpHeaders({
        'Authorization': 'Bearer ' + CredentialControlService.getToken(),
        "Content-Type": "application/json"
      })
    };

    const body = {
      'nombre':nombre,
      'url':url,
      'categoria':categoria
    }

    return this.http.post(this.urlCreateVideo, body, httpOptions);
  }

  deleteVideo(id:string){

    const httpOptions = {
      headers: new HttpHeaders({
        'Authorization': 'Bearer ' + CredentialControlService.getToken(),
      })
    };

    return this.http.delete(this.urlDeleteVideo+id,httpOptions);
  }

  editVideo(id:string, nombre:string, url:string, categoria:string){
    const httpOptions = {
      headers: new HttpHeaders({
        'Authorization': 'Bearer ' + CredentialControlService.getToken(),
        "Content-Type": "application/json"
      })
    };

    const body = {
      'nombre':nombre,
      'url':url,
      'categoria':categoria
    }

    return this.http.put(this.urlEditVideo+id, body, httpOptions);
  }
}
