import { UpperCasePipe } from '@angular/common';
import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { DomSanitizer } from '@angular/platform-browser';
import { CategoryCrudServiceService } from '../common-services/category-crud-service.service';
import { CredentialControlService } from '../common-services/credential-control.service';
import { VideoCrudServiceService } from '../common-services/video-crud-service.service';
import { Category } from '../interfaces/category';
import { VideosInterface } from '../interfaces/videos-interface';

@Component({
  selector: 'app-start',
  templateUrl: './start.component.html',
  styleUrls: ['./start.component.css'],
  encapsulation: ViewEncapsulation.None
})

export class StartComponent implements OnInit {

  total:number;
  videoList:Array<VideosInterface> = [];
  filteredVideoList:Array<VideosInterface> = [];
  categoryList:Array<Category> = [];
  filteredCategoryList:Array<string> = [];
  filteredVideosByCategory:Array<VideosInterface> = [];
  filteredVideosBySearch:Array<VideosInterface> = [];
  isAdmin:boolean = false;
  theme:number = 1;

  constructor(
    private cred: CredentialControlService, 
    private videoCrud: VideoCrudServiceService, 
    private catCrud: CategoryCrudServiceService, 
    public sanitizer:DomSanitizer) 
    { 
    this.total = 0;
  }

  ngOnInit(): void {
    this.cred.checkCredential().subscribe(
      res => {
        this.total = res['total'];
        console.log('Token is still valid. Total: ' + this.total);
        this.getVideoList();
        this.getCategoryList();
      },
      err => {
        console.log(err);
        localStorage.clear;
        location.assign('/Login')
      }
    );
    if(localStorage.getItem('Role')?.includes('ADMIN')) 
      this.isAdmin = true;
  }

  private getVideoList(){
    this.videoCrud.getVideos(this.total,0).subscribe(
      res => {
        this.videoList = res['productos'];
        this.filteredVideoList = this.videoList;
        this.filteredVideosByCategory = this.videoList
        this.filteredVideosBySearch = this.videoList
        console.log(res);
      },
      err => {
        console.log(err);
      }
    )
  }

  private getCategoryList(){
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

  logOut(){
    this.cred.logOut();
  }

  filterSearch(data: string) {
    this.filteredVideosBySearch = this.videoList.filter(
      x => x.nombre.includes(
        data.toUpperCase()
      )
    );
    this.filteredVideoList = this.inBoth(this.filteredVideosByCategory, this.filteredVideosBySearch);
  }

  filterCategories(data: string, checked: boolean) {

    if (!checked) {
      this.filteredCategoryList.push(data)
    }
    else {
      this.filteredCategoryList = this.filteredCategoryList.filter(x => x != data);
    }

    console.log(checked);
    console.log(this.filteredCategoryList);

    this.filteredVideosByCategory = this.videoList

    this.filteredCategoryList.forEach(
      x => this.filteredVideosByCategory = this.filteredVideosByCategory.filter(
        y => !y.categoria.nombre.includes(x)
      )
    );

    this.filteredVideoList = this.inBoth(this.filteredVideosByCategory, this.filteredVideosBySearch);
  }

  //theming
  themeChange(){    
    switch(this.theme){
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
 
  operation = (list1:Array<VideosInterface>, list2:Array<VideosInterface>, isUnion = false) =>
  list1.filter( a => isUnion === list2.some( b => a._id === b._id ) );

  inBoth = (list1:Array<VideosInterface>, list2:Array<VideosInterface>) => this.operation(list1, list2, true);
}
