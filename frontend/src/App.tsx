import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { CampaignListPage } from './pages/CampaignListPage';
import { CampaignCreatePage } from './pages/CampaignCreatePage';
import { CampaignDetailPage } from './pages/CampaignDetailPage';
import './App.css';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<CampaignListPage />} />
        <Route path="/campaigns/new" element={<CampaignCreatePage />} />
        <Route path="/campaigns/:id" element={<CampaignDetailPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
