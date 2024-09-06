import { SVGProps } from "react";

export type IconSvgProps = SVGProps<SVGSVGElement> & {
  size?: number;
};

export type File = {
  id: string;
  name: string;
  path: string;
  type: string;
  created_at: string;
  last_update: string;
};
