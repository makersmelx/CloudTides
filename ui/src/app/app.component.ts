import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { PRODUCT_NAME } from '@tide-config/const';

import { LoginService, UserInfo } from './login/login.service';
import { TranslateService } from '@ngx-translate/core';
import { I18nService } from './config/i18n';

@Component({
  selector: 'tide-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {

  constructor(
    private readonly loginService: LoginService,
    private readonly router: Router,
    translate: TranslateService,
    i18nService: I18nService,
  ) {
    translate.setDefaultLang('en');
    const language = i18nService.getLanguage();
    translate.use(language);
  }

  readonly vo = {
    title: PRODUCT_NAME,
    user$: this.loginService.session$,
  };

  signOut() {
    this.loginService.logout().subscribe();
  }

  inLoginPage() {
    return this.router.url.indexOf('/login') === 0;
  }

  ngOnInit() {

  }
}
