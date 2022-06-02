import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { CredentialControlService } from '../common-services/credential-control.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class LoginComponent implements OnInit {

  constructor(private cred:CredentialControlService) {  }

  onSubmit(uname: string, pwd: string) {
    console.log('Done');
    this.cred.sendCredential(uname, pwd).subscribe(
        res => {
            if (res.hasOwnProperty('token')) {
              localStorage.setItem('Token', res['token']);
              localStorage.setItem('Role', res['usuario'].rol);
              location.assign('/Start')
            }
        },
        err => {
            console.log(err);
        }
    );
  }

  ngOnInit(): void {
    console.log('Current token is: ' + CredentialControlService.getToken());
    this.cred.checkCredential().subscribe(
      res => {
        location.assign('/Start');
      },
      err => {
        localStorage.clear();
      }
    )
  }
}
