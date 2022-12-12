import { Route, Routes } from 'solid-app-router';
import type { Component } from 'solid-js';

import styles from './App.module.scss';
import Menu from './Menu';

const App: Component = () => {
  return (

    <div class={`container-lg ${styles.App}`}>
      <Menu />
      <Routes>
        <Route path='/' element={<div>root</div>}   />
        <Route path="/:test" element={<div>test</div>} />
      </Routes>
    </div>
  );
};

export default App;
