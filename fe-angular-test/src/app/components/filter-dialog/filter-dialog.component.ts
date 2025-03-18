import { ChangeDetectionStrategy, Component } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { provideNativeDateAdapter } from '@angular/material/core';
import { MatButtonModule } from '@angular/material/button';
import { FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { MatDialogModule } from '@angular/material/dialog';

@Component({
  selector: 'app-filter-dialog',
  templateUrl: './filter-dialog.component.html',
  providers: [provideNativeDateAdapter()],
  imports: [
    MatFormFieldModule,
    MatDialogModule,
    MatInputModule,
    MatSelectModule,
    MatDatepickerModule,
    MatButtonModule,
    ReactiveFormsModule
  ],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class FilterDialogComponent {
  filterForm: FormGroup;

  constructor(
    private dialogRef: MatDialogRef<FilterDialogComponent>,
    private fb: FormBuilder
  ) {
    this.filterForm = this.fb.group({
      dob: [null],
      status: ['All']
    });
  }

  onCancel(): void {
    this.dialogRef.close();
  }

  onApply(): void {
    this.dialogRef.close(this.filterForm.value);
  }
}
