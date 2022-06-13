import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { DomSanitizer } from '@angular/platform-browser';
import { CategoryCrudServiceService } from '../common-services/category-crud-service.service';
import { CredentialControlService } from '../common-services/credential-control.service';
import { UserCrudService } from '../common-services/user-crud.service';
import { VideoCrudServiceService } from '../common-services/video-crud-service.service';
import { Category } from '../interfaces/category';
import { User } from '../interfaces/user-interface';
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
  filteredVideoList: Array<VideosInterface> = [];
  filteredCategoryList: Array<Category> = [];
  filteredUserList: Array<User> = [];
  totalVideos: number = 0;
  theme: number = 1;
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
        this.filteredVideoList = this.videoList;
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
        this.filteredCategoryList = this.categoryList;
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
        this.filteredUserList = this.userList;
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

  //Delete storage and return to login
  logOut() {
    this.cred.logOut();
  }

  //Set active section in admin panel
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

  //Set active parameters in edit
  setEditUser(id: string, nameElement: HTMLInputElement, emailElement: HTMLInputElement) {
    let selectedUser = this.userList.filter(x => x.uid.includes(id))[0];
    nameElement.value = selectedUser.nombre;
    emailElement.value = selectedUser.correo
  }

  setEditVideo(id: string, nameElement: HTMLInputElement, urlElement: HTMLInputElement, categoryElement: HTMLSelectElement) {
    let selectedVideo = this.videoList.filter(x => x._id.includes(id))[0];
    nameElement.value = selectedVideo.nombre;
    urlElement.value = selectedVideo.url;
    categoryElement.value = selectedVideo.categoria._id;
  }

  setEditCategory(id: string, nameElement: HTMLInputElement) {
    let selectedCategory = this.categoryList.filter(x => x._id.includes(id))[0];
    nameElement.value = selectedCategory.nombre;
  }

  //Searchbox utilities
  filterData(searchValue: string) {
    this.filteredCategoryList = this.categoryList.filter(x => x.nombre.toUpperCase().includes(searchValue.toUpperCase()));
    this.filteredUserList = this.userList.filter(x => x.nombre.toUpperCase().includes(searchValue.toUpperCase()));
    this.filteredVideoList = this.videoList.filter(x => x.nombre.toUpperCase().includes(searchValue.toUpperCase()));
  }

  //theming
  themeChange() {
    switch (this.theme) {
      case 1:
        document.getElementById('tema')?.setAttribute('src', 'assets/img/Giroro.jpeg')
        document.body.classList.toggle("redtema");
        break;
      case 2:
        document.getElementById('tema')?.setAttribute('src', 'assets/img/Kururu.jpeg')
        document.body.classList.toggle("redtema");
        document.body.classList.toggle("yellowtema");
        break;
      case 3:
        document.getElementById('tema')?.setAttribute('src', 'assets/img/Avatar.jpeg')
        document.body.classList.toggle("yellowtema");
        this.theme = 0;
        break;
    }
    this.theme += 1;
  }

  //Get video per category
  categoryCount(id: string) {
    console.log();
    
    return this.videoList.filter(x => 
      { 
        let works : boolean;
        try {
          works = x.categoria._id.includes(id);
        }
        catch {
          works =  x.categoria._id == id;
        }
        return works;
      }).length;
  }

  /*
    Video Data Manipulation
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

  sendEditCategory(id: string, nombre: string) {
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

  sendDeleteCategory(id: string, value: string) {
    if (value.includes('0')) {
      if (confirm('¿Estás seguro de que deseas eliminar la categoría?'))
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
    } else alert('No se puede eliminar una categoría con videos');

  }

  sendCreateCategory(nombre: string) {
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

  /*
    User Data Manipulation
  */

  sendEditUser(id: string, nombre: string, correo: string, password: string) {
    this.userCrud.editUser(id, nombre, correo, password).subscribe(
      res => {
        alert('Se ha modificado la información satisfactoriamente.');
        this.getUserList();
        console.log(res);
      },
      err => {
        alert('Connection failed. Check console log for details.');
        console.log(err);
      }
    )
  }

  sendDeleteUser(id: string) {
    if (confirm('¿Estás seguro de que deseas eliminar el usuario?'))
      this.userCrud.deleteUser(id).subscribe(
        res => {
          alert('Se ha eliminado satisfactoriamente.');
          this.getUserList();
          console.log(res);
        },
        err => {
          alert('Connection failed. Check console log for details.');
          console.log(err);
        }
      )
  }

  sendCreateUser(nombre: string, correo: string, password: string) {
    this.userCrud.createUser(nombre, correo, password).subscribe(
      res => {
        alert('Se ha modificado la información satisfactoriamente.');
        this.getUserList();
        console.log(res);
      },
      err => {
        alert('Connection failed. Check console log for details.');
        console.log(err);
      }
    )
  }
}
