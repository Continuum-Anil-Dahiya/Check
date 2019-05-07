
USE Nocbo

GO
IF NOT EXISTS(SELECT 1 FROM sys.columns
WHERE Name = N'Origin' AND OBJECT_ID = OBJECT_ID(N'[dbo].[jmgttasknotesclosed]'))
BEGIN
ALTER TABLE jmgttasknotesclosed ADD Origin VARCHAR(25)
END

GO
IF NOT EXISTS(SELECT 1 FROM sys.columns
WHERE Name = N'Origin' AND OBJECT_ID = OBJECT_ID(N'[dbo].[jmgttask_noc_notes_CLOSED]'))
BEGIN
ALTER TABLE jmgttask_noc_notes_CLOSED ADD Origin VARCHAR(25)
END
