import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { DomSanitizer } from '@angular/platform-browser';
import { CategoryCrudServiceService } from '../common-services/category-crud-service.service';
import { CredentialControlService } from '../common-services/credential-control.service';
import { UserCrudService } from '../common-services/user-crud.service';
import { VideoCrudServiceService } from '../common-services/video-crud-service.service';
import { Category } from '../interfaces/category';
import { User } from '../interfaces/user-interface';
import { Users } from '../interfaces/users';
import { VideosInterface } from '../interfaces/videos-interface';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class AdminComponent implements OnInit {

  videoList: Array<VideosInterface> = [];
  categoryList: Array<Category> = [];
  userList: Array<User> = [];
  totalVideos: number = 0;

  constructor(private cred: CredentialControlService,
    private videoCrud: VideoCrudServiceService,
    private catCrud: CategoryCrudServiceService,
    private userCrud: UserCrudService,
    public sanitizer: DomSanitizer) {
    this.totalVideos = 0;
  }

  ngOnInit(): void {
    const body = document.querySelector('body'),
      sidebar = body?.querySelector('nav'),
      toggle = body?.querySelector(".toggle"),
      searchBtn = body?.querySelector(".search-box"),
      modeText = body?.querySelector(".mode-text");

    toggle?.addEventListener("click", () => {
      sidebar?.classList.toggle("close");
    })

    searchBtn?.addEventListener("click", () => {
      sidebar?.classList.remove("close");
    })

    this.cred.checkCredential().subscribe(
      res => {
        this.totalVideos = res['total'];
        console.log('Token is still valid.');
        this.getVideoList();
        this.getCategoryList();
        this.getUserList();
      },
      err => {
        console.log(err);
        localStorage.clear;
        location.assign('/Login')
      }
    );
  }

  /*
    Get all the elements and store them
  */

  private getVideoList() {
    this.videoCrud.getVideos(this.totalVideos, 0).subscribe(
      res => {
        this.videoList = res['productos'];
        console.log(res);
      },
      err => {
        console.log(err);
      }
    )
  }

  private getCategoryList() {
    this.catCrud.getCategories().subscribe(
      res => {
        this.categoryList = res['categorias'];
        console.log(res);
      },
      err => {
        console.log(err);
      }
    )
  }

  private getUserList() {
    this.userCrud.getUsers().subscribe(
      res => {
        this.userList = res['usuarios'];
        console.log(res);
      },
      err => {
        console.log(err);
      }
    )
  }

  /*
    Active function for click events
  */

  logOut() {
    this.cred.logOut();
  }

  setActiveSection(tab: number) {

    let sections = document.querySelectorAll('section');
    sections.forEach(el => {
      el.setAttribute('hidden', 'true')
    });

    switch (tab) {
      case 1:
        document.getElementById("videos")?.toggleAttribute('hidden');
        break;
      case 2:
        document.getElementById("categoria")?.toggleAttribute('hidden');
        break;
      case 3:
        document.getElementById("usuario")?.toggleAttribute('hidden');
        break;
      case 4:
        document.getElementById("modificar")?.toggleAttribute('hidden');
        break;
      case 5:
        document.getElementById("crear")?.toggleAttribute('hidden');
        break;
    }
  }

  /*
    Video data manipulation
  */
  sendCreateVideo(nombre: string, url: string, categoria: string) {

    this.videoCrud.createVideo(nombre, url, categoria).subscribe(
      res => {
        console.log(res);
        this.totalVideos += 1;
        this.getVideoList();
        alert('Video uploaded')
      },
      err => {
        alert('Connection failed. Check console log for details.')
        console.log(err);
      }
    )
  }

  sendDeleteVideo(id: string) {
    if (confirm('¿Estás seguro de que deseas eliminar el video?'))
      this.videoCrud.deleteVideo(id).subscribe(
        res => {
          console.log(res);
          this.videoList = this.videoList.filter(x => (!x._id.includes(id)));
        },
        err => {
          alert('Connection failed. Check console log for details.');
          console.log(err);
        }
      );
  }

  sendEditVideo(id: string, nombre: string, url: string, categoria: string) {
    this.videoCrud.editVideo(id, nombre, url, categoria).subscribe(
      res => {
        alert('Se ha modificado la información satisfactoriamente.');
        this.getVideoList();
        console.log(res);
      },
      err => {
        alert('Connection failed. Check console log for details.');
        console.log(err);
      }
    )
  }

  /*
    Category Data Manipulation
  */

  sendEditCategory(id: string, nombre: string){
    this.catCrud.editCategory(id, nombre).subscribe(
      res => {
        alert('Se ha modificado la información satisfactoriamente.');
        this.getCategoryList();
        console.log(res);
      },
      err => {
        alert('Connection failed. Check console log for details.');
        console.log(err);
      }
    )
  }

  sendDeleteCategory(id: string){
    this.catCrud.deleteCategory(id).subscribe(
      res => {
        alert('Se ha eliminado satisfactoriamente.');
        this.getCategoryList();
        console.log(res);
      },
      err => {
        alert('Connection failed. Check console log for details.');
        console.log(err);
      }
    )
  }

  sendCreateCategory(nombre: string){
    this.catCrud.createCategory(nombre).subscribe(
      res => {
        alert('Se ha modificado la información satisfactoriamente.');
        this.getCategoryList();
        console.log(res);
      },
      err => {
        alert('Connection failed. Check console log for details.');
        console.log(err);
      }
    )
  }
}
