import { Fragment } from 'react';
import PropTypes from 'prop-types';

import styles from './Template.module.css';

import Icon from 'components/Form/Icon'
import Label from 'components/Form/Label'
import Button from 'components/Form/Button'

export const SIDE_VISIBILITY = {
  AUTO: "auto",
  SHOW: "show",
  HIDE: "hide"
}

export const Template = ({ 
  side, templateKey, dirty, className,
  getText, onEvent,
  ...props 
}) => {
  const itemMenu = (keyValue, classValue, eventValue, labelValue) => {
    return <Button id={keyValue} key={keyValue}
      className={classValue}
      onClick={ ()=>onEvent(...eventValue) }
      value={labelValue}
    />
  }
  return(
    <div {...props}
      className={`${styles.sidebar} ${((side !== "auto")? side : "")} ${className}`} >
      {itemMenu("cmd_back",
        `${"medium"} ${styles.itemButton} ${styles.selected}`, 
        ["checkTemplate",[{}, "LOAD_SETTING"]],
        <Label value={getText("label_back")} 
          leftIcon={<Icon iconKey="Reply" />} iconWidth="25px"  />
      )}
      <div key="tmp_sep_1" className={styles.separator} />

      {(!["_blank", "_sample"].includes(templateKey))?
      <Fragment>
        <div key="tmp_sep_2" className={styles.separator} />
        {itemMenu("cmd_save",
          `${"full medium"} ${styles.itemButton} ${(dirty)?styles.selected:""}`,
          ["saveTemplate", [true]],
          <Label value={getText("template_save")} 
            leftIcon={<Icon iconKey="Check" />} iconWidth="25px"  />
        )}
        {itemMenu("cmd_create",
          `${"full medium"} ${styles.itemButton}`,
          ["createTemplate", []],
          <Label value={getText("template_create_from")} 
            leftIcon={<Icon iconKey="Sitemap" />} iconWidth="25px" />
        )}
        {itemMenu("cmd_delete",
          `${"full medium"} ${styles.itemButton}`,
          ["deleteTemplate", []],
          <Label value={getText("label_delete")} 
            leftIcon={<Icon iconKey="Times" />} iconWidth="20px"  />
        )}
      </Fragment>:null}

      <div key="tmp_sep_3" className={styles.separator} />
      {itemMenu("cmd_blank",
        `${"full medium"} ${styles.itemButton}`,
        ["checkTemplate", [{}, 'NEW_BLANK']],
        <Label value={getText("template_new_blank")} 
          leftIcon={<Icon iconKey="Plus" />} iconWidth="25px" />
      )}
      {itemMenu("cmd_sample",
        `${"full medium"} ${styles.itemButton}`,
        ["checkTemplate", [{}, 'NEW_SAMPLE']],
        <Label value={getText("template_new_sample")} 
          leftIcon={<Icon iconKey="Plus" />} iconWidth="25px" />
      )}

      <div key="tmp_sep_4" className={styles.separator} />
      {itemMenu("cmd_print",
        `${"full medium"} ${styles.itemButton}`, 
        ["showPreview", []],
        <Label value={getText("label_print")} 
          leftIcon={<Icon iconKey="Print" />} iconWidth="25px"  />
      )}
      {itemMenu("cmd_json",
        `${"full medium"} ${styles.itemButton}`,
        ["exportTemplate", []],
        <Label value={getText("template_export_json")} 
          leftIcon={<Icon iconKey="Code" />} iconWidth="25px" />
      )}

      <div key="tmp_sep_5" className={styles.separator} />
      {itemMenu("cmd_help",
        `${"full medium"} ${styles.itemButton}`, 
        ["showHelp", ["program/editor"]],
        <Label value={getText("label_help")} 
          leftIcon={<Icon iconKey="QuestionCircle" />} iconWidth="20px"  />
      )}

    </div>
  )
  
}

Template.propTypes = {
  /**
   * SideBar visibility
   */
  side: PropTypes.oneOf(Object.values(SIDE_VISIBILITY)).isRequired,
  templateKey: PropTypes.string.isRequired,
  dirty: PropTypes.bool,
  className: PropTypes.string, 
  /**
   * Menu selection handle
   */
  onEvent: PropTypes.func,
  /**
   * Localization
   */
  getText: PropTypes.func,
}

Template.defaultProps = {
  side: SIDE_VISIBILITY.AUTO,
  templateKey: "",
  dirty: false, 
  className: "",  
  onEvent: undefined,
  getText: undefined,
}

export default Template;