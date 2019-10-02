import { Component, OnInit } from '@angular/core';
import{File, FileService} from'../file.service';
@Component({
  selector: 'app-file',
  templateUrl: './file.component.html',
  styleUrls: ['./file.component.scss']
})
export class FileComponent implements OnInit {

  files: File[];

  constructor(private fileService: FileService) { }

  ngOnInit() {
    this.loadAll();
  }
  loadAll() {
    this.fileService.getFileList().subscribe((data:File[])=> this.files = data);
  }

}
