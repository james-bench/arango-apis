//
// DISCLAIMER
//
// Copyright 2020 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//
// Author Robert Stam
//

package v1

// SpecEquals returns true when source & other have the same specification values
// Note that several fields, like ID, Name and Description are not checked
func (source *BackupPolicy) SpecEquals(other *BackupPolicy) bool {
	return source.GetEmailNotification() == other.GetEmailNotification() &&
		source.GetIsDeleted() == other.GetIsDeleted() &&
		source.GetIsPaused() == other.GetIsPaused() &&
		source.GetUpload() == other.GetUpload() &&
		source.GetSchedule().Equals(other.GetSchedule()) &&
		source.GetRetentionPeriod().Equal(other.GetRetentionPeriod())
}

// Equals returns true when source & other have the same values
func (source *BackupPolicy_Schedule) Equals(other *BackupPolicy_Schedule) bool {
	return source.GetScheduleType() == other.GetScheduleType() &&
		source.GetHourlySchedule().Equals(other.GetHourlySchedule()) &&
		source.GetDailySchedule().Equals(other.GetDailySchedule()) &&
		source.GetMonthlySchedule().Equals(other.GetMonthlySchedule())
}

// Equals returns true when source & other have the same values
func (source *BackupPolicy_HourlySchedule) Equals(other *BackupPolicy_HourlySchedule) bool {
	return source.GetScheduleEveryIntervalHours() == other.GetScheduleEveryIntervalHours()
}

// Equals returns true when source & other have the same values
func (source *BackupPolicy_DailySchedule) Equals(other *BackupPolicy_DailySchedule) bool {
	return source.GetMonday() == other.GetMonday() &&
		source.GetTuesday() == other.GetTuesday() &&
		source.GetWednesday() == other.GetWednesday() &&
		source.GetThursday() == other.GetThursday() &&
		source.GetFriday() == other.GetFriday() &&
		source.GetSaturday() == other.GetSaturday() &&
		source.GetSunday() == other.GetSunday() &&
		source.GetScheduleAt().Equals(other.GetScheduleAt())
}

// Equals returns true when source & other have the same values
func (source *BackupPolicy_MonthlySchedule) Equals(other *BackupPolicy_MonthlySchedule) bool {
	return source.GetDayOfMonth() == other.GetDayOfMonth() &&
		source.GetScheduleAt().Equals(other.GetScheduleAt())
}

// Equals returns true when source & other have the same values
func (source *BackupPolicy_Status) Equals(other *BackupPolicy_Status) bool {
	return source.GetMessage() == other.GetMessage() &&
		source.GetNextBackup().Equal(other.GetNextBackup())
}

// Equals returns true when source & other have the same values
func (source *TimeOfDay) Equals(other *TimeOfDay) bool {
	return source.GetHours() == other.GetHours() &&
		source.GetMinutes() == other.GetMinutes() &&
		source.GetTimeZone() == other.GetTimeZone()
}

// SpecEquals returns true when source & other have the same specification values
// Note that several fields, like ID, Name and Description are not checked
func (source *Backup) SpecEquals(other *Backup) bool {
	return source.GetAutoDeletedAt().Equal(other.GetAutoDeletedAt()) &&
		source.GetIsDeleted() == other.GetIsDeleted() &&
		source.GetUpload() == other.GetUpload() &&
		source.GetDownload().Equals(other.GetDownload())
}

// Equals returns true when source & other have the same values
func (source *Backup_DownloadSpec) Equals(other *Backup_DownloadSpec) bool {
	return source.GetRevision() == other.GetRevision() &&
		source.GetLastUpdatedAt().Equal(other.GetLastUpdatedAt())
}

// Equals returns true when source & other have the same values
func (source *Backup_Status) Equals(other *Backup_Status) bool {
	return source.GetMessage() == other.GetMessage() &&
		source.GetAvailable() == other.GetAvailable() &&
		source.GetIsFailed() == other.GetIsFailed() &&
		source.GetCreatedAt().Equal(other.GetCreatedAt()) &&
		source.GetProgress() == other.GetProgress() &&
		source.GetState() == other.GetState() &&
		source.GetVersion() == other.GetVersion() &&
		source.GetSizeBytes() == other.GetSizeBytes() &&
		source.GetDbservers() == other.GetDbservers() &&
		source.GetUploadStatus().Equals(other.GetUploadStatus()) &&
		source.GetDownloadStatus().Equals(other.GetDownloadStatus()) &&
		source.GetUploadOnly() == other.GetUploadOnly()
}

// Equals returns true when source & other have the same values
func (source *Backup_UploadStatus) Equals(other *Backup_UploadStatus) bool {
	return source.GetUploaded() == other.GetUploaded() &&
		source.GetUploadedAt().Equal(other.GetUploadedAt()) &&
		source.GetSizeBytes() == other.GetSizeBytes()
}

// Equals returns true when source & other have the same values
func (source *Backup_DownloadStatus) Equals(other *Backup_DownloadStatus) bool {
	return source.GetDownloaded() == other.GetDownloaded() &&
		source.GetDownloadedAt().Equal(other.GetDownloadedAt()) &&
		source.GetRevision() == other.GetRevision()
}
