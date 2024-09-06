"use client";
import {
  ArrowDownTrayIcon,
  EllipsisVerticalIcon,
  PencilIcon,
  TrashIcon,
} from "@heroicons/react/24/outline";
import { Button } from "@nextui-org/button";
import {
  Dropdown,
  DropdownItem,
  DropdownMenu,
  DropdownTrigger,
} from "@nextui-org/dropdown";
import * as React from "react";
import { useRef } from "react";
import { useDisclosure } from "@nextui-org/modal";
import { toast } from "react-toastify";

import { RenameFileModal } from "./rename-file-modal";

import { File } from "@/types";

export interface IFileActionsMenuProps {
  file: File;
  onAction: () => Promise<void>;
}

export function FileActionsMenu(props: IFileActionsMenuProps) {
  const renameFileModalDiscosure = useDisclosure();
  const downloadLinkRef = useRef<HTMLAnchorElement>(null);

  const deleteFile = async () => {
    try {
      await fetch(
        `${process.env.NEXT_PUBLIC_VFS_BASE_URL}/${props.file.type}${props.file.path}`,
        { method: "DELETE" },
      );
    } catch (error) {
      toast("Fail to delete files", { type: "error" });
    } finally {
      await props.onAction();

      return;
    }
  };

  const downloadFile = async (e: React.MouseEvent) => {
    e.stopPropagation();
    try {
      const response = await fetch(
        `${process.env.NEXT_PUBLIC_VFS_BASE_URL}/file/content${props.file.path}`,
      );

      if (!response.ok) {
        throw new Error("fail to download document");
      }

      const blob = await response.blob();
      const url = window.URL.createObjectURL(blob);

      if (!downloadLinkRef.current) {
        throw new Error("invalid download link html element");
      }
      downloadLinkRef.current.href = url;
      downloadLinkRef.current.download = props.file.name;
      downloadLinkRef.current.click();
      window.URL.revokeObjectURL(url);
    } catch (error) {
      toast("Fail to download files", { type: "error" });
    }
  };

  const openModal = (e: React.MouseEvent) => {
    e.stopPropagation();
    renameFileModalDiscosure.onOpen();
  };

  return (
    <div>
      <Dropdown>
        <DropdownTrigger>
          <Button
            isIconOnly
            variant="light"
            onClick={(e: React.MouseEvent) => {
              e.stopPropagation();
            }}
          >
            <EllipsisVerticalIcon className="size-5" />
          </Button>
        </DropdownTrigger>
        {props.file.type === "file" && (
          <DropdownMenu>
            <DropdownItem
              startContent={<ArrowDownTrayIcon className="size-5" />}
              onClick={downloadFile}
            >
              Download
              <a ref={downloadLinkRef} className="hidden" href="/">
                downlonad
              </a>
            </DropdownItem>
            <DropdownItem
              showDivider
              startContent={<PencilIcon className="size-5" />}
              onClick={openModal}
            >
              Rename
            </DropdownItem>
            <DropdownItem
              className="text-danger"
              color={"danger"}
              startContent={<TrashIcon className="size-5" />}
              onClick={deleteFile}
            >
              Delete
            </DropdownItem>
          </DropdownMenu>
        )}

        {props.file.type === "directory" && (
          <DropdownMenu>
            <DropdownItem
              showDivider
              startContent={<PencilIcon className="size-5" />}
              onClick={openModal}
            >
              Rename
            </DropdownItem>
            <DropdownItem
              className="text-danger"
              color={"danger"}
              startContent={<TrashIcon className="size-5" />}
              onClick={deleteFile}
            >
              Delete
            </DropdownItem>
          </DropdownMenu>
        )}
      </Dropdown>
      <RenameFileModal
        disclosure={renameFileModalDiscosure}
        file={props.file}
        onFileRenamed={props.onAction}
      />
    </div>
  );
}
