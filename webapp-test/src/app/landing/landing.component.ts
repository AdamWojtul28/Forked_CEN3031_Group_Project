import { Component } from '@angular/core';

@Component({
  selector: 'app-landing',
  templateUrl: './landing.component.html',
  styleUrls: ['./landing.component.css']
})
export class LandingComponent {
  login = {
    email: 'a@gmail.com',
    password: '12345678'
  };
  status = 0;

  constructor() {}
  ngOnInit() {
    this.postMethod();
  }

  /*
  Endpoint: 'https://localhost:9000/api/user/signin'
  Placeholder: 'https://jsonplaceholder.typicode.com/posts'
  */
  postMethod() {
    fetch('https://localhost:9000/api/user/signin', {
      method: 'POST',
      body: JSON.stringify(this.login),
      headers: {
        'Content-type': 'application/json; charset=UTF-8',
      },
    })
    .then((response) => {
      console.log(response.type);
      if(response.status === 200) {
        console.log('Status 200');
      }
      else if(response.status == 201) {
        console.log('Status 201');
      }
      else {
        console.log('Error');
      }
    })
    //.then((json) => this.login = json);
    .catch(error => {console.log('Fetch failed')})
  }
}