import { Component, OnInit } from '@angular/core';
import { DomSanitizer } from '@angular/platform-browser';
import { MatSnackBar } from '@angular/material/snack-bar';

import { of } from 'rxjs/observable/of';
import { catchError, switchMap } from 'rxjs/operators';
import { MainService, Stats } from '../../services/main.service';

function ImageURL(rollnum: string, userid: string) {
    const iitkhome = `https://home.iitk.ac.in/~${ userid }/dp`;
    const pcluboaimage = `https://photos.pclub.in/oaphoto/${ rollnum }_0.jpg`;
    return `url("${ iitkhome }"), url("${ pcluboaimage }")`;
}

@Component({
  selector: 'puppy-results',
  templateUrl: './results.component.html',
  styleUrls: [ './results.component.scss' ]
})
export class ResultsComponent implements OnInit {

  user$: any;
  matches: any[];
  stats: Stats;

  colorScheme = {
    domain: ['#C2024F', '#04BBBF', '#D2D945', '#FCB13F', '#FF594F']
  };

  constructor(private main: MainService,
              private sanitizer: DomSanitizer,
              private snackbar: MatSnackBar) {}

  ngOnInit() {
    this.user$ = this.main.user$;
    this.doSubmit();
    this.getStats();
  }

  get url() {
    const currentUser = {
      ...this.main.user$.value
    };
    return this.sanitizer.bypassSecurityTrustStyle(ImageURL(currentUser._id, currentUser.email));
  }

  get registrations() {
    const stats = this.stats;
    const totalMales = stats.othermales + stats.y18males + stats.y19males + stats.y20males + stats.y21males;
    const totalFemales = stats.otherfemales + stats.y18females + stats.y19females + stats.y20females + stats.y21females;
    return [{
      name: 'Males',
      value: totalMales,
    }, {
      name: 'Females',
      value: totalFemales,
    }];
  }

  get hearts() {
    const stats = this.stats;
    const totalMaleHearts = stats.othermaleHearts + stats.y18maleHearts + stats.y19maleHearts + stats.y20maleHearts + stats.y21maleHearts;
    const totalFemaleHearts = stats.otherfemaleHearts + stats.y18femaleHearts + stats.y19femaleHearts + stats.y20femaleHearts + stats.y21femaleHearts;
    return [{
      name: 'Males',
      value: totalMaleHearts,
    }, {
      name: 'Females',
      value: totalFemaleHearts,
    }];
  }
  get fhearts() {
    const stats = this.stats;
    return [{
      name: 'Others',
      value: stats.otherfemaleHearts,
    }, {
      name: 'Y18',
      value: stats.y18femaleHearts,
    }, {
      name: 'Y19',
      value: stats.y19femaleHearts,
    }, {
      name: 'Y20',
      value: stats.y20femaleHearts,
    }, {
        name: 'Y21',
        value: stats.y21femaleHearts,
    }].reverse();

  }

  get mhearts() {
    const stats = this.stats;
    return [{
      name: 'Others',
      value: stats.othermaleHearts,
    }, {
      name: 'Y18',
      value: stats.y18maleHearts,
    }, {
      name: 'Y19',
      value: stats.y19maleHearts,
    }, {
      name: 'Y20',
      value: stats.y20maleHearts,
    }, {
        name: 'Y21',
        value: stats.y21maleHearts,
    }].reverse();
  }

  get fregs() {
    const stats = this.stats;
    return [{
      name: 'Others',
      value: stats.otherfemales,
    }, {
      name: 'Y18',
      value: stats.y18females,
    }, {
      name: 'Y19',
      value: stats.y19females,
    }, {
      name: 'Y20',
      value: stats.y20females,
    }, {
        name: 'Y21',
        value: stats.y21females,
    }].reverse();

  }

  get mregs() {
    const stats = this.stats;
    return [{
      name: 'Others',
      value: stats.othermales,
    }, {
      name: 'Y18',
      value: stats.y18males,
    }, {
      name: 'Y19',
      value: stats.y19males,
    }, {
      name: 'Y20',
      value: stats.y20males,
    }, {
        name: 'Y21',
        value: stats.y21males,
    }].reverse();
  }

  maleHearts(user) {
    return user.data.received.filter((x) => x.genderOfSender === '1');
  }

  femaleHearts(user) {
    return user.data.received.filter((x) => x.genderOfSender === '0');
  }

  doSubmit() {
    const user = this.user$.value;
    this.main.submit().pipe(
      catchError((err) => of(console.error(err))),
      switchMap(() => this.main.matches()),
    ).subscribe(
      (match) => {
        if (match.matches === '') {
          this.matches = [];
        } else {
          this.matches = match.matches.split(' ').map(x => this.main.people.filter(p => p._id === x)[0]);
        }
      },
      (error) => this.snackbar.open(error, '', { duration: 3000 })
    );
  }

  getStats() {
    this.main.stats().subscribe((x) => this.stats = x);
  }

  onLogout() {
    location.reload();
  }
}
