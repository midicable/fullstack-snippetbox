import { AppLayout } from "app/layout";
import { SnippetPage } from "pages/SnippetPage";
import { SnippetsPage } from "pages/SnippetsPage";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { AppRoute } from "shared";

export const AppRouter = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path={AppRoute.ROOT} element={<AppLayout />}>
          <Route index element={<div>main</div>} />
          <Route path={AppRoute.SNIPPETS} element={<SnippetsPage />} />
          <Route path={AppRoute.SNIPPET} element={<SnippetPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
};
