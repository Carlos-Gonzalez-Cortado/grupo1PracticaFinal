import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { DataType } from '../data-type';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class LoginComponent implements OnInit {

  constructor(private http: HttpClient) { 

  }

  sendCredential(email: string, password: string) {
    const url = 'https://labinfsoft.herokuapp.com/api/auth/login';
    const body_content = {
        "correo": email, 
        "password": password
    }

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type':  'application/json'
      })
    };    

    return this.http.post(url, body_content, httpOptions) as Observable<DataType>;
  }

  onSubmit(uname: string, pwd: string) {
    console.log('Done');
    this.sendCredential(uname, pwd).subscribe(
        res => {
            console.log(res);
            
            if (res.hasOwnProperty('token')) {
              localStorage.setItem('Token', res['token']);
            }
        },
        err => {
            console.log(err);
        }
    );
}

  myFunc(){
    console.log("function called");
  }

  ngOnInit(): void {

  }



}
