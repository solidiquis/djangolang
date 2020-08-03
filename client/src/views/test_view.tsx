import * as React from 'react';
import * as ReactDOM from 'react-dom';

import { AsyncButton } from '../components/async_button';

export const TestView: React.FC = () => {
  return (
    <div>
      <AsyncButton text="Hello World" color="red" endpoint="/api" />
      <AsyncButton text="Foo" color="red" endpoint="/api/foo" />
      <AsyncButton text="Bar" color="red" endpoint="/api/bar" />
    </div>
  );
};
