import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { GetVideoInterface } from '../interfaces/get-video-interface';
import { CredentialControlService } from './credential-control.service';

@Injectable({
  providedIn: 'root'
})
export class VideoCrudServiceService {

  url = 'https://labinfsoft.herokuapp.com';
  urlGetVideos = this.url + '/api/videos';

  constructor(private http: HttpClient) { }

  getVideos(limite:number, desde:number){

    const fullyQualifiedUrl = this.urlGetVideos + '?limite=' + limite + '&desde=' + desde;

    console.log(fullyQualifiedUrl);
    

    const httpOptions = {
      headers: new HttpHeaders({
        'Authorization': 'Bearer ' + CredentialControlService.getToken()
      })
    };

    return this.http.get(this.urlGetVideos, httpOptions) as Observable<GetVideoInterface>;
  }
}
