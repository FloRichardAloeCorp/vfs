import path from "path";

import { Button } from "@nextui-org/button";
import {
  Dropdown,
  DropdownItem,
  DropdownMenu,
  DropdownSection,
  DropdownTrigger,
} from "@nextui-org/dropdown";
import { Input } from "@nextui-org/input";
import * as React from "react";
import { useRef, useState } from "react";
import { toast } from "react-toastify";

export interface INewFolderDropdownProps {
  isOpen?: boolean;
  onNewFolder: (newPath: string) => void;
  currentPath: string;
}

export function NewFolderDropdown(props: INewFolderDropdownProps) {
  const [dropdownOpen, setDropdownOpen] = useState(props.isOpen);
  const folderNameInputRef = useRef<HTMLInputElement>(null);

  const newFolder = async (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key != "Enter") {
      return;
    }

    const folderPath = path.join(
      props.currentPath,
      folderNameInputRef.current?.value as string,
    );

    try {
      await fetch(
        `${process.env.NEXT_PUBLIC_VFS_BASE_URL}/directory${folderPath}`,
        { method: "POST" },
      );
    } catch (error) {
      toast("Fail to create folder", { type: "error" });
    } finally {
      setDropdownOpen(false);
      props.onNewFolder(folderPath);
    }
  };

  return (
    <Dropdown
      closeOnSelect={false}
      isOpen={dropdownOpen}
      onClose={() => setDropdownOpen(false)}
    >
      <DropdownTrigger>
        <Button
          className="min-h-0 min-w-0 h-fit px-0"
          variant="light"
          onClick={() => setDropdownOpen(true)}
        >
          New folder
        </Button>
      </DropdownTrigger>
      <DropdownMenu variant="light">
        <DropdownSection title="Folder name">
          <DropdownItem
            description="Press enter to validate"
            onClick={() => {
              folderNameInputRef.current?.focus();
            }}
          >
            <Input
              ref={folderNameInputRef}
              className="pb-2"
              type="text"
              onKeyDown={newFolder}
            />
          </DropdownItem>
        </DropdownSection>
      </DropdownMenu>
    </Dropdown>
  );
}
