// src/components/Table.tsx
import React from 'react';
import './Table.css';

const Table: React.FC = () => {
    return (
        <table className="ticket-table">
            <thead>
                <tr>
                    <th>Showtime</th>
                    <th>Check in time</th>
                    <th>ticket_id</th>
                    <th>Status</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td>10:00 AM</td>
                    <td>9:45 AM</td>
                    <td>12345</td>
                    <td>Checked In</td>
                </tr>
                <tr>
                    <td>12:00 PM</td>
                    <td>11:45 AM</td>
                    <td>12346</td>
                    <td>Checked In</td>
                </tr>
                <tr>
                    <td>2:00 PM</td>
                    <td>1:45 PM</td>
                    <td>12347</td>
                    <td>Expire</td>
                </tr>
                <tr>
                    <td>4:00 PM</td>
                    <td>3:45 PM</td>
                    <td>12348</td>
                    <td>Checked In</td>
                </tr>
            </tbody>
        </table>
    );
};

export default Table;
