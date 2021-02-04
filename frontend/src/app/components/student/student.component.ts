import { Component, EventEmitter, Input, Output } from '@angular/core';
import { DomSanitizer } from '@angular/platform-browser';

function ImageURL(rollnum: string, userid: string) {
    const iitkhome = `https://home.iitk.ac.in/~${ userid }/dp`;
    const pcluboaimage = `https://search.pclub.in/oaphoto/${ rollnum }_0.jpg`;
    const oaimage = `https://oa.cc.iitk.ac.in/Oa/Jsp/Photo/${ rollnum }_0.jpg`;
    return `url("${ iitkhome }"), url("${ pcluboaimage }"), url("${ oaimage }")`;
}

@Component({
  selector: 'puppy-student',
  templateUrl: './student.component.html',
  styleUrls: ['./student.component.css'],
})
export class StudentComponent {

  @Input()
  student: any;
  @Output()
  select = new EventEmitter();

  constructor(private sanitizer: DomSanitizer) {}

  get url() {
    return this.sanitizer.bypassSecurityTrustStyle(ImageURL(this.student._id, this.student.email));
  }

  selectStudent() {
    this.select.emit();
  }

}
