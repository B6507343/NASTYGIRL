import React, { useState } from "react";
import { QrReader } from "react-qr-reader";
import { Form, Input, Button,message } from "antd"; // Import Ant Design components
import { PlusOutlined } from "@ant-design/icons"; // Import the icon
import "./QrScanner.css"; // Import the CSS file
import{Checkin} from "../../service/https/ticketcheck";
import { TicketcheckInterface } from "../../interface/ITicketcheck.ts";
import { TicketInterface } from "../../interface/ITicket.ts";


const QrScanner: React.FC = () => {
  const [data, setData] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [form] = Form.useForm(); // Create a form instance


  const handleError = (err: any) => {
    console.error(err);
    if (!data) {
      setError("Error scanning QR code");
    }
  };

  const handleScan = (scannedData: string | null) => {
    console.log("Scanned Data:", scannedData);
    if (scannedData) {
      setData(scannedData);
      setError(null);
    } else if (!data) {
      setError("No data found in QR code");
    }
  };

  
  // const newTicketcheck={
  //   TicketID : Number(data),
  //   TimeStamp: new Date().toISOString(),
  //   Status: "Ok",
  //   Member: 1
    
  // }

  const handleSubmit = () => {
    
    if (Number(data)) {  
      
      console.log("Data submitted:", data);
      message.success("Data submitted successfully!");
      Checkin(Number(data));

     
    } else {
      message.error("No data to submit.");
    }
  };
 
  
  

  return (
    <div className="container">
      <div className="scannerContainer">
        <h1>QR Code Scanner</h1>
        <QrReader
          onResult={(result, error) => {
            if (result) {
              const text = result.getText ? result.getText() : result.text || "";
              handleScan(text);
            }

            if (error) {
              handleError(error);
            }
          }}
          className="qrReader" // Apply the CSS class
        />

        {/* Ant Design Form */}
        <Form
          form={form}
          layout="vertical"
          initialValues={{ ticketID: data }} // Set initial value to scanned data
        >
          <div className="input-button-container"> {/* Flexbox container */}
            <Form.Item
              label="Ticket ID"
              name="ticketID"
              rules={[
                {
                  required: true,
                  message: "Please Enter Ticket ID",
                },
              ]}
              style={{ flex: 1 }} // Make input take available space
            >
              <Input
                placeholder="Please Enter Ticket ID"
                allowClear
                onChange={(e) => setData(e.target.value)} // Update state with input value
              />
            </Form.Item>

            <Button type="primary" onClick={handleSubmit}  icon={<PlusOutlined />} style={{ marginLeft: '10px' }}>
              Enter
            </Button>
          </div>
        </Form>
      </div>
    </div>
  );
};

export default QrScanner;

