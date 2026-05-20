import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { CampaignListPage } from './pages/CampaignListPage';
import { CampaignCreatePage } from './pages/CampaignCreatePage';
import './App.css';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<CampaignListPage />} />
        <Route path="/campaigns/new" element={<CampaignCreatePage />} />
        <Route
          path="/campaigns/:id"
          element={
            <div className="page-container">
              <p className="placeholder-message">
                Campaign detail page will be implemented later.
              </p>
            </div>
          }
        />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
