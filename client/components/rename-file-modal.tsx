import type { UseDisclosureReturn } from "@nextui-org/use-disclosure";

import { Button } from "@nextui-org/button";
import { Input } from "@nextui-org/input";
import {
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
} from "@nextui-org/modal";
import * as React from "react";
import { useState } from "react";
import { toast } from "react-toastify";

import { File } from "@/types";

export interface IRenameFileModalProps {
  disclosure: UseDisclosureReturn;
  file: File;
  onFileRenamed: () => Promise<void>;
}

export function RenameFileModal(props: IRenameFileModalProps) {
  const { isOpen, onOpenChange } = props.disclosure;
  const [newFileName, setNewFileName] = useState("");

  const renameFile = (onClose: () => void) => {
    return async () => {
      try {
        await fetch(
          `${process.env.NEXT_PUBLIC_VFS_BASE_URL}/${props.file.type}/name${props.file.path}?name=${newFileName}`,
          {
            method: "PUT",
          },
        );
        props.onFileRenamed();
      } catch (error) {
        toast("Fail to rename file", { type: "error" });
      }
      onClose();
    };
  };

  return (
    <Modal isOpen={isOpen} onOpenChange={onOpenChange}>
      <ModalContent>
        {(onClose) => (
          <>
            <ModalHeader>Rename file</ModalHeader>
            <ModalBody>
              <Input
                type="text"
                value={newFileName}
                onClick={(e: React.MouseEvent) => {
                  e.stopPropagation();
                }}
                onValueChange={setNewFileName}
              />
            </ModalBody>

            <ModalFooter>
              <Button color="danger" variant="light" onPress={onClose}>
                Close
              </Button>
              <Button color="primary" onPress={renameFile(onClose)}>
                Rename
              </Button>
            </ModalFooter>
          </>
        )}
      </ModalContent>
    </Modal>
  );
}
