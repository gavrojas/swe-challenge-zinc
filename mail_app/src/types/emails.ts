export interface EmailDocument {
  message_id: string;
  date: string;
  from: string;
  to: string;
  subject: string;
  mime_version: string;
  content_type: string;
  content_transfer_encoding: string;
  x_from: string;
  x_to: string;
  x_cc: string;
  x_bcc: string;
  x_folder: string;
  x_origin: string;
  x_filename: string;
  body: string;
  user: string;
  folder_path: string;
}