import * as React from 'react';
import axios from 'axios';

type SomeProp = {
  text: string;
  color: string;
  endpoint: string;
};

export const AsyncButton: React.FC<SomeProp> = ({
  text,
  color,
  endpoint
}: SomeProp) => {
  const ping = async (e) => {
    try {
      const res = await axios.get(endpoint);
      console.log(res);
    } catch(e) {
      console.log(e);
    };
  };

  return(
    <button
      onClick={ping}
      style={{color: color}}
    >{text}
    </button>
  );
};
