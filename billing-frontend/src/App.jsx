import { useEffect, useState } from "react";

function App() {
  const [locations, setLocations] = useState([]);
  const [consumers, setConsumers] = useState([]);
  const [selectedLocation, setSelectedLocation] = useState("");
  const [selectedConsumer, setSelectedConsumer] = useState("");
  const [date, setDate] = useState("");
  const [bill, setBill] = useState(null);

  // Load locations
  useEffect(() => {
    fetch("http://localhost:8080/locations")
      .then(res => res.json())
      .then(data => setLocations(data));
  }, []);

  // Load consumers when location changes
  useEffect(() => {
    if (selectedLocation) {
      fetch(`http://localhost:8080/consumers?location_id=${selectedLocation}`)
        .then(res => res.json())
        .then(data => setConsumers(data));
    }
  }, [selectedLocation]);

 const getBill = () => {
  if (!selectedConsumer || !date) {
    alert("Please select consumer and date");
    return;
  }

  // Ensure correct format (safety)
  const formattedDate = new Date(date).toISOString().split("T")[0];

  fetch(`http://localhost:8080/bill?consumer_id=${selectedConsumer}&date=${formattedDate}`)
    .then(res => res.json())
    .then(data => setBill(data));
};

  return (
    <div style={{ padding: "20px" }}>
      <h2>⚡ Electricity Billing System</h2>

      <select onChange={(e) => setSelectedLocation(e.target.value)}>
        <option>Select Location</option>
        {locations.map(l => (
          <option key={l.ID} value={l.ID}>{l.Name}</option>
        ))}
      </select>

      <br /><br />

      <select onChange={(e) => setSelectedConsumer(e.target.value)}>
        <option>Select Consumer</option>
        {consumers.map(c => (
          <option key={c.ID} value={c.ID}>{c.Name}</option>
        ))}
      </select>

      <br /><br />

      <input type="date" onChange={(e) => setDate(e.target.value)} />

      <br /><br />

      <button onClick={getBill}>Get Bill</button>

{bill && (
  <div style={{ marginTop: "20px" }}>
    <h3>Total: ₹ {bill.total}</h3>

    {/* ✅ NEW ADDITIONS */}
    <p><b>First Reading:</b> {bill.first_reading}</p>
    <p><b>Last Reading:</b> {bill.last_reading}</p>
    <p><b>Units Consumed:</b> {bill.last_reading - bill.first_reading}</p>

    <hr />

    {bill.breakdown.map((b, i) => (
      <div key={i}>
        Units: {b.units} | Rate: {b.rate} | Amount: {b.amount}
      </div>
    ))}
  </div>
)}
    </div>
  );
}

export default App;