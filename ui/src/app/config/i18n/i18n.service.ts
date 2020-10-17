import { Injectable } from '@angular/core';

@Injectable()
export class I18nService {
  constructor() {
    if (!localStorage.getItem('i18n')) {
      localStorage.setItem('i18n', 'en');
    }
  }

  getLanguage(): string {
    return localStorage.getItem('i18n');
  }

  setLanguage(newLanguage: string): any {
    localStorage.setItem('i18n', newLanguage);
  }
}
