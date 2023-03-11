import { Route, Routes } from 'solid-app-router';
import type { Component } from 'solid-js';
import AboutView from './AboutView';

import BlogListView from './BlogListView';
import BlogView from './BlogView';
import Menu from './Menu';

const App: Component = () => {
  return (

    <div class={`container-lg`}>
      <Menu />
      <Routes>
        <Route path='/' element={<BlogListView />}   />
        <Route path='/post/:id' element={<BlogView />}   />
        <Route path="/about" element={<AboutView />} />
      </Routes>
    </div>
  );
};

export default App;
