"use client";

import path from "path";

import { DocumentPlusIcon, FolderPlusIcon } from "@heroicons/react/24/outline";
import { Input } from "@nextui-org/input";
import { Listbox, ListboxItem } from "@nextui-org/listbox";
import * as React from "react";
import { ChangeEvent, useRef, useState } from "react";
import { toast } from "react-toastify";

import { NewFolderDropdown } from "./new-folder-dropdown";

import { usePath } from "@/context/path";

export interface ISideMenuProps {
  onNewFile: () => void;
  onNewFolder: (newPath: string) => void;
}

export function SideMenu(props: ISideMenuProps) {
  const { currentPath } = usePath();
  const [dropdownOpen, setDropdownOpen] = useState(false);

  const fileInputRef = useRef<HTMLInputElement>(null);

  const handleNewFileClick = (_e: React.MouseEvent<HTMLLIElement>) => {
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  const newFile = async (e: ChangeEvent<HTMLInputElement>) => {
    if (e.target.files == null || e.target.files.length === 0) {
      return;
    }

    const file = e.target.files[0];

    await fetch(
      `${process.env.NEXT_PUBLIC_VFS_BASE_URL}/file/${path.join(currentPath, file.name)}`,
      {
        method: "POST",
        body: file,
      },
    ).catch((_error) => {
      toast("Fail to create file", { type: "error" });
    });
    props.onNewFile();
  };

  return (
    <div className="border-small rounded-small mx-5 my-4">
      <Input
        ref={fileInputRef}
        className="hidden"
        type="file"
        onChange={newFile}
      />
      <Listbox>
        <ListboxItem
          key="new_file"
          startContent={<DocumentPlusIcon className="size-5" />}
          onClick={handleNewFileClick}
        >
          {" "}
          New File{" "}
        </ListboxItem>
        <ListboxItem
          key="new_folder"
          startContent={<FolderPlusIcon className="size-5" />}
          textValue="test"
          onClick={() => setDropdownOpen(true)}
        >
          <NewFolderDropdown
            currentPath={currentPath}
            isOpen={dropdownOpen}
            onNewFolder={props.onNewFolder}
          />
        </ListboxItem>
      </Listbox>
    </div>
  );
}
