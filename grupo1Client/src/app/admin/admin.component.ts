import { Component, OnInit, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class AdminComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
    const body = document.querySelector('body'),
      sidebar = body?.querySelector('nav'),
      toggle = body?.querySelector(".toggle"),
      searchBtn = body?.querySelector(".search-box"),
      modeText = body?.querySelector(".mode-text");


        toggle?.addEventListener("click" , () =>{
            sidebar?.classList.toggle("close");
        })

        searchBtn?.addEventListener("click" , () =>{
            sidebar?.classList.remove("close");
        })
  }

}
