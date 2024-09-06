"use client";

import { createContext, useContext, useState } from "react";

const PathContext = createContext({
  currentPath: "",
  setCurrentPath: (_val: string) => {},
});

const PathProvider = ({ children }: { children: React.ReactNode }) => {
  const [currentPath, setCurrentPath] = useState("/");

  return (
    <PathContext.Provider value={{ currentPath, setCurrentPath }}>
      {children}
    </PathContext.Provider>
  );
};

const usePath = () => {
  return useContext(PathContext);
};

export { PathProvider, usePath };
