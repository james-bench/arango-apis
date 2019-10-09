//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Robert Stam
//

package v1

import (
	"fmt"

	common "github.com/arangodb-managed/apis/common/v1"
)

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

// Validate returns an error if not correct, otherwise nil
func (source *BackupPolicy_Schedule) Validate() error {
	if source == nil {
		return common.InvalidArgument("source not set")
	}
	hourly := false
	daily := false
	monthly := false
	switch source.GetScheduleType() {
	case "":
		return common.InvalidArgument("schedule type not set")
	case "Hourly":
		hourly = true
		if err := source.GetHourlySchedule().Validate("HourlySchedule"); err != nil {
			return err
		}
	case "Daily":
		daily = true
		if err := source.GetDailySchedule().Validate("DailySchedule"); err != nil {
			return err
		}
	case "Monthly":
		monthly = true
		if err := source.GetMonthlySchedule().Validate("MonthlySchedule"); err != nil {
			return err
		}
	default:
		return common.InvalidArgument("schedule type '%s' unknown", source.GetScheduleType())
	}
	if !hourly && source.GetHourlySchedule() != nil {
		return common.InvalidArgument("not allowed to set field '%s' in type %s", "HourlySchedule", source.GetScheduleType())
	}
	if !daily && source.GetDailySchedule() != nil {
		return common.InvalidArgument("not allowed to set field '%s' in type %s", "DailySchedule", source.GetScheduleType())
	}
	if !monthly && source.GetMonthlySchedule() != nil {
		return common.InvalidArgument("not allowed to set field '%s' in type %s", "MonthlySchedule", source.GetScheduleType())
	}
	return nil
}

// Validate returns an error if not correct, otherwise nil
func (source *BackupPolicy_HourlySchedule) Validate(prefix string) error {
	if source == nil {
		return common.InvalidArgument("required field '%s' not found", prefix)
	}
	// Validate ScheduleEveryIntervalHours
	if source.GetScheduleEveryIntervalHours() < 1 || source.GetScheduleEveryIntervalHours() > 23 {
		return common.InvalidArgument("field '%s.%s' should beween 1-23 (both included)", prefix, "ScheduleEveryIntervalHours")
	}
	return nil
}

// Validate returns an error if not correct, otherwise nil
func (source *BackupPolicy_DailySchedule) Validate(prefix string) error {
	if source == nil {
		return common.InvalidArgument("required field '%s' not found", prefix)
	}
	// Day bools are OK by design, Validate ScheduleAt
	if err := source.GetScheduleAt().Validate(fmt.Sprintf("%s.ScheduleAt", prefix)); err != nil {
		return err
	}
	return nil
}

// Validate returns an error if not correct, otherwise nil
func (source *BackupPolicy_MonthlySchedule) Validate(prefix string) error {
	if source == nil {
		return common.InvalidArgument("required field '%s' not found", prefix)
	}
	// Validate ScheduleEveryIntervalHours
	if source.GetDayOfMonth() < 1 || source.GetDayOfMonth() > 31 {
		return common.InvalidArgument("field '%s.%s' should beween 1-31 (both included)", prefix, "DayOfMonth")
	}
	if err := source.GetScheduleAt().Validate(fmt.Sprintf("%s.ScheduleAt", prefix)); err != nil {
		return err
	}
	return nil
}

// Validate returns an error if not correct, otherwise nil
func (source *TimeOfDay) Validate(prefix string) error {
	if source == nil {
		return common.InvalidArgument("required field '%s' not found", prefix)
	}
	if source.GetHours() < 0 || source.GetHours() > 23 {
		return common.InvalidArgument("field '%s.%s' should beween 0-23 (both included)", prefix, "Hours")
	}
	if source.GetMinutes() < 0 || source.GetMinutes() > 59 {
		return common.InvalidArgument("field '%s.%s' should beween 0-59 (both included)", prefix, "Minutes")
	}
	return nil
}
