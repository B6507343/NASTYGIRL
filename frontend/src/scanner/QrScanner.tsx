import React, { useState } from "react";
import { QrReader } from "react-qr-reader";
import { Form, Input, Button, message } from "antd"; // Import Ant Design components
import { PlusOutlined } from "@ant-design/icons"; // Import the icon
import "./QrScanner.css"; // Import the CSS file
import { Checkin } from "../../service/https/ticketcheck"; // Import the service that calls the backend

const QrScanner: React.FC = () => {
  const [data, setData] = useState<string | null>(null);
  const [form] = Form.useForm(); // Create a form instance

  const handleSubmit = async () => {
    if (Number(data)) {
      try {
        // ส่ง ticket_id ไปยัง backend ผ่าน Checkin function
        const response = await Checkin(Number(data));
        if (response.status) {
          message.success("Data submitted successfully!");
          form.resetFields(); // ล้างข้อมูลหลังจาก submit สำเร็จ
        } else {
          message.error("Failed to submit data.");
        }
      } catch (error) {
        message.error("Error occurred while submitting data.");
      }
    } else {
      message.error("No data to submit.");
    }
  };

  return (
    <div className="container">
      <div className="scannerContainer">
        <h1>QR Code Scanner</h1>
        <QrReader
          onResult={(result: any, error: any) => {
            if (result) {
              const text = result?.text || result.getText?.() || "";
              if (text) {
                setData(text);
              }
            }
            if (error) {
              console.error(error);
            }
          }}
          constraints={{ facingMode: "environment" }}
          className="qrReader"
        />

        {/* Ant Design Form */}
        <Form form={form} layout="vertical">
          <div className="input-button-container">
            <Form.Item
              label="Ticket ID"
              name="ticketID"
              rules={[
                {
                  required: true,
                  message: "Please Enter Ticket ID",
                },
              ]}
              style={{ flex: 1 }}
            >
             <Input
  placeholder="Please Enter Ticket ID"
  allowClear
  value={data || ""}
  onChange={(e: React.ChangeEvent<HTMLInputElement>) => setData(e.target.value)} // ระบุชนิดของ e ที่นี่
/>

            </Form.Item>

            <Button
              type="primary"
              onClick={handleSubmit}
              icon={<PlusOutlined />}
              style={{ marginLeft: "10px" }}
            >
              Enter
            </Button>
          </div>
        </Form>
      </div>
    </div>
  );
};

export default QrScanner;
