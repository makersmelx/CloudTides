import { Injectable } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';

@Injectable()
export class I18nService {
  constructor(
    public readonly translate: TranslateService,
  ) {
    this.resetLanguage();
    translate.use(localStorage.getItem('i18n'));
  }

  resetLanguage(): any {
    if (!localStorage.getItem('i18n')) {
      localStorage.setItem('i18n', 'en');
    }
  }

  getLanguage(): string {
    this.resetLanguage();
    return localStorage.getItem('i18n');
  }

  setLanguage(newLanguage: string): any {
    localStorage.setItem('i18n', newLanguage);
    this.translate.use(newLanguage);
  }
}
