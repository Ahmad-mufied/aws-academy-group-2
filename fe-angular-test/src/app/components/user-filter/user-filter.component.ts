import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { CommonModule } from '@angular/common';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { FormsModule } from '@angular/forms';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatNativeDateModule } from '@angular/material/core';
import { MatSelectModule } from '@angular/material/select';
import { MatDialogModule } from '@angular/material/dialog';
import { ChangeDetectionStrategy } from '@angular/core';
import { provideNativeDateAdapter } from '@angular/material/core';

@Component({
  selector: 'app-user-filter',
  standalone: true,
  imports: [
    CommonModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    FormsModule,
    MatDatepickerModule,
    MatNativeDateModule,
    MatSelectModule,
    MatDialogModule
  ],
  providers: [provideNativeDateAdapter()],
  templateUrl: './user-filter.component.html',
  styleUrls: ['./user-filter.component.css'],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class UserFilterComponent {
  filters = {
    startDate: null as Date | null,
    endDate: null as Date | null,
    status: ''
  };

  constructor(
    public dialogRef: MatDialogRef<UserFilterComponent>,
    @Inject(MAT_DIALOG_DATA) public data: { startDate?: Date; endDate?: Date; status?: string }
  ) {
    if (data) {
      this.filters = {
        startDate: data.startDate || null,
        endDate: data.endDate || null,
        status: data.status || ''
      };
    }
  }

  onCancel(): void {
    this.dialogRef.close();
  }

  onApply(): void {
    this.dialogRef.close(this.filters);
  }
}